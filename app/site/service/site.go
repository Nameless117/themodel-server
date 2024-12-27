package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/site/models"
	"gorm.io/gorm"
	"strings"
	"sync"
)

type WebSite struct {
	service.Service
}

// Get 获取数据
func (w *WebSite) Get() (*SiteData, error) {
	var (
		wg     sync.WaitGroup
		errCh  = make(chan error, 6) // 错误通道
		result = &SiteData{}         // 初始化返回结果
	)

	// 创建一个互斥锁来保护并发写入
	var mu sync.Mutex

	// Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		config, err := w.GetConfig()
		if err != nil {
			errCh <- fmt.Errorf("get config error: %w", err)
			return
		}
		if config != nil {
			mu.Lock()
			result.LogoLarge = config.LogoLarge
			result.LogoSmall = config.LogoSmall
			result.BannerBgImg = config.BannerBgImg
			result.BannerBook = config.BannerBook
			mu.Unlock()
		}
	}()

	// Banner
	wg.Add(1)
	go func() {
		defer wg.Done()
		banner, err := w.GetBanner()
		if err != nil {
			errCh <- fmt.Errorf("get banner error: %w", err)
			return
		}

		if banner != nil {
			if banner.HomeBanner != nil {
				mu.Lock()
				result.HomeBanner.Desc = banner.HomeBanner.Description
				mu.Unlock()
			}
			mu.Lock()
			result.AudioBanner = AudioBanner{
				Desc:     banner.AudioBanner.Description,
				SubDesc:  banner.AudioBanner.SubDescription,
				PDFTitle: banner.AudioBanner.PdfTitle,
				PDFURL:   banner.AudioBanner.PdfUrl,
			}
			mu.Unlock()
		}
	}()

	// Footer
	wg.Add(1)
	go func() {
		defer wg.Done()
		footer, err := w.GetFooter()
		if err != nil {
			errCh <- fmt.Errorf("get footer error: %w", err)
			return
		}
		if footer != nil {
			mu.Lock()
			result.Footer = Footer{
				Title:    footer.Title,
				Email:    footer.Email,
				SubTitle: footer.SubTitle,
				Text:     footer.Text,
			}
			mu.Unlock()
		}
	}()

	// MenuItem
	wg.Add(1)
	go func() {
		defer wg.Done()
		menuItems, err := w.GetMenuItems()
		if err != nil {
			errCh <- fmt.Errorf("get menu item error: %w", err)
			return
		}
		if len(menuItems) != 0 {
			mu.Lock()
			var menuList []MenuItem
			for _, item := range menuItems {
				menuList = append(menuList, MenuItem{
					Title: item.Title,
					Path:  item.Path,
				})
			}
			result.MenuList = menuList
			mu.Unlock()
		}

	}()

	// SubConfig
	wg.Add(1)
	go func() {
		defer wg.Done()
		subConfig, err := w.GetSubConfig()
		if err != nil {
			errCh <- fmt.Errorf("get sub config error: %w", err)
			return
		}
		if subConfig != nil {
			mu.Lock()
			result.Subscription.Title = subConfig.Title
			result.Subscription.ExampleDesc = subConfig.ExampleDesc
			result.Subscription.ExampleURL = subConfig.ExampleUrl
			mu.Unlock()
		}

	}()

	// Links
	wg.Add(1)
	go func() {
		defer wg.Done()
		links, err := w.GetLinks()
		if err != nil {
			errCh <- fmt.Errorf("get links error: %w", err)
			return
		}
		if len(links) != 0 {
			mu.Lock()
			var subscriptionList []SubscriptionItem
			for _, link := range links {
				subscriptionList = append(subscriptionList, SubscriptionItem{
					Title: link.Title,
					URL:   link.Url,
				})
			}
			result.Subscription.List = subscriptionList
			mu.Unlock()
		}

	}()

	// 等待所有goroutine完成
	wg.Wait()

	// 关闭错误通道
	close(errCh)

	// 检查是否有错误发生
	if len(errCh) > 0 {
		// 收集所有错误
		var errMsgs []string
		for err := range errCh {
			errMsgs = append(errMsgs, err.Error())
		}
		return result, fmt.Errorf("multiple errors occurred: %s", strings.Join(errMsgs, "; "))
	}

	return result, nil
}

// GetConfig 获取SysSiteConfig对象
func (w *WebSite) GetConfig() (*models.SysSiteConfig, error) {
	var data models.SysSiteConfig
	err := w.Orm.Model(&models.SysSiteConfig{}).Scopes().First(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在")
		return nil, err
	}
	if err != nil {
		w.Log.Errorf("db error:%s", err)
		return nil, err
	}
	return &data, nil
}

// GetBanner 获取SysSiteBanners对象
// BannerResult 用于存储两种类型的 banner
type BannerResult struct {
	HomeBanner  *models.SysSiteBanners `json:"homeBanner"`
	AudioBanner *models.SysSiteBanners `json:"audioBanner"`
}

// GetBanner 获取首页和有声读物的 banner
func (w *WebSite) GetBanner() (*BannerResult, error) {
	result := &BannerResult{}

	// 使用事务确保数据一致性
	err := w.Orm.Transaction(func(tx *gorm.DB) error {
		// 获取 home 类型的 banner
		homeBanner := &models.SysSiteBanners{}
		if err := tx.Where("type = ?", "home").First(homeBanner).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				w.Log.Warn("home banner not found")
				// 继续执行，不返回错误
			} else {
				return fmt.Errorf("query home banner error: %w", err)
			}
		} else {
			result.HomeBanner = homeBanner
		}

		// 获取 audio 类型的 banner
		audioBanner := &models.SysSiteBanners{}
		if err := tx.Where("type = ?", "audio").First(audioBanner).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				w.Log.Warn("audio banner not found")
				// 继续执行，不返回错误
			} else {
				return fmt.Errorf("query audio banner error: %w", err)
			}
		} else {
			result.AudioBanner = audioBanner
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 检查是否至少有一个 banner 存在
	if result.HomeBanner == nil && result.AudioBanner == nil {
		return nil, errors.New("no banner found")
	}

	return result, nil
}

// GetFooter 获取SysSiteFooterConfig对象
func (w *WebSite) GetFooter() (*models.SysSiteFooterConfig, error) {
	var data models.SysSiteFooterConfig

	err := w.Orm.Model(&models.SysSiteFooterConfig{}).Scopes().First(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在")
		return nil, err
	}
	if err != nil {
		w.Log.Errorf("db error:%s", err)
		return nil, err
	}
	return &data, nil
}

// GetMenuItems 获取SysSiteMenuItems对象
func (w *WebSite) GetMenuItems() ([]*models.SysSiteMenuItems, error) {
	var list []*models.SysSiteMenuItems

	err := w.Orm.Model(&models.SysSiteMenuItems{}).Scopes().Order("sortOrder ASC").Find(&list).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在")
		return nil, err
	}
	if err != nil {
		w.Log.Errorf("db error:%s", err)
		return nil, err
	}
	return list, nil
}

// GetSubConfig 获取SysSiteSubscriptionConfig对象
func (w *WebSite) GetSubConfig() (*models.SysSiteSubscriptionConfig, error) {
	var data models.SysSiteSubscriptionConfig

	err := w.Orm.Model(&models.SysSiteSubscriptionConfig{}).Scopes().First(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在")
		return nil, err
	}
	if err != nil {
		w.Log.Errorf("db error:%s", err)
		return nil, err
	}
	return &data, nil
}

// GetLinks 获取SysSiteSubscriptionLinks对象
func (w *WebSite) GetLinks() ([]*models.SysSiteSubscriptionLinks, error) {
	var list []*models.SysSiteSubscriptionLinks

	err := w.Orm.Model(&models.SysSiteSubscriptionLinks{}).Scopes().Order("sortOrder ASC").Find(&list).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在")
		return nil, err
	}
	if err != nil {
		w.Log.Errorf("db error:%s", err)
		return nil, err
	}
	return list, nil
}

// MenuItem 表示菜单项
type MenuItem struct {
	Title string `json:"title"`
	Path  string `json:"path"`
}

// SubscriptionItem 表示订阅项
type SubscriptionItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Subscription 表示订阅信息
type Subscription struct {
	Title       string             `json:"title"`
	List        []SubscriptionItem `json:"list"`
	ExampleDesc string             `json:"exampleDesc"`
	ExampleURL  string             `json:"exampleUrl"`
}

// AudioBanner 表示有声读物 Banner 信息
type AudioBanner struct {
	Desc     string `json:"desc"`
	SubDesc  string `json:"subDesc"`
	PDFTitle string `json:"PDFTitle"`
	PDFURL   string `json:"PDFUrl"`
}

// Banner 表示首页 Banner 信息
type Banner struct {
	Desc string `json:"desc"`
}

// Footer 表示页脚信息
type Footer struct {
	Title    string `json:"title"`
	Email    string `json:"email"`
	SubTitle string `json:"subTitle"`
	Text     string `json:"text"`
}

// SiteData 是顶层结构，表示所有的配置数据
type SiteData struct {
	LogoLarge    string       `json:"logoLarge"`
	LogoSmall    string       `json:"logoSmall"`
	BannerBgImg  string       `json:"bannerBgImg"`
	BannerBook   string       `json:"bannerBook"`
	MenuList     []MenuItem   `json:"menuList"`
	HomeBanner   Banner       `json:"homeBanner"`
	AudioBanner  AudioBanner  `json:"audioBanner"`
	Subscription Subscription `json:"subscription"`
	Footer       Footer       `json:"footer"`
}
