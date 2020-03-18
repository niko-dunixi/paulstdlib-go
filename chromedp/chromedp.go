package chromedp

import (
	"context"
	"github.com/chromedp/chromedp"
)

// Literally the same thing as chromedp.DefaultExecAllocatorOptions. The only
// difference is that I've skipped the headless option on purpose 99.9% of the
// time while I'm building new tools, I want to see what chromedp is doing.
func VisibleWithDefaultOptions() []chromedp.ExecAllocatorOption {
	allocatorOptions := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	return allocatorOptions
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
