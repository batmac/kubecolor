package printer

import (
	"io"

	"github.com/dty1er/kubecolor/color"
	"github.com/dty1er/kubecolor/kubectl"
)

type KubectlOutputColoredPrinter struct {
	SubcommandInfo *kubectl.SubcommandInfo
	DarkBackground bool
}

func (kp *KubectlOutputColoredPrinter) Print(r io.Reader, w io.Writer) {
	withHeader := !kp.SubcommandInfo.NoHeader

	var printer Printer = &SingleColoredPrinter{Color: color.Green} // default in green

	switch kp.SubcommandInfo.Subcommand {
	case kubectl.Top, kubectl.APIResources:
		printer = &TablePrinter{WithHeader: withHeader, DarkBackground: kp.DarkBackground}

	case kubectl.Get:
		switch {
		case kp.SubcommandInfo.FormatOption == kubectl.None:
			printer = &TablePrinter{WithHeader: withHeader, DarkBackground: kp.DarkBackground}
		default:
			printer = &GetPrinter{WithHeader: withHeader, FormatOpt: kp.SubcommandInfo.FormatOption, DarkBackground: kp.DarkBackground}
		}

	case kubectl.Describe:
		printer = &DescribePrinter{DarkBackground: kp.DarkBackground}
	}

	printer.Print(r, w)
}
