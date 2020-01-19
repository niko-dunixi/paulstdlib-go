package chromedp

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

// Literally the same thing as chromedp.DefaultExecAllocatorOptions. The only
// difference is that I've skipped the headless option on purpose 99.9% of the
// time while I'm building new tools, I want to see what chromedp is doing.
func VisibleWithDefaultOptions() []chromedp.ExecAllocatorOption {
	return []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		func(allocator *chromedp.ExecAllocator) {
			//chromedp.Flag("headless", true)(allocator)
			chromedp.Flag("hide-scrollbars", true)(allocator)
			chromedp.Flag("mute-audio", true)(allocator)
		},
		// After Puppeteer's default behavior.
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-features", "site-per-process,TranslateUI,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", true),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),
	}
}

func VisibleChromedp(ctx context.Context, opts ...chromedp.ContextOption) (context.Context, context.CancelFunc) {
	allocator, allocCancelFunc := chromedp.NewExecAllocator(ctx, VisibleWithDefaultOptions()...)
	chromeCtx, chromeCtxCancelFunc := chromedp.NewContext(allocator, opts...)
	// Wrap both cancel functions and then call them LIFO order
	dualCancelFunc := func() {
		chromeCtxCancelFunc()
		allocCancelFunc()
	}
	return chromeCtx, dualCancelFunc
}
