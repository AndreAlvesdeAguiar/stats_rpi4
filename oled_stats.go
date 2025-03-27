package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/waxdred/go-i2c-oled"
	"github.com/waxdred/go-i2c-oled/ssd1306"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Stat struct {
	Ip   string
	Cpu  string
	Mem  string
	Disk string
}

func executeCmd(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return ""
	}
	return strings.TrimSpace(out.String())
}

func GetStat() Stat {
	return Stat{
		Ip:   executeCmd("bash", "-c", "hostname -I | cut -d' ' -f1"),
		Cpu:  executeCmd("bash", "-c", "top -bn1 | grep load | awk '{printf \"CPU Load: %.2f\", $(NF-2)}'"),
		Mem:  executeCmd("bash", "-c", "free -m | awk 'NR==2{printf \"Mem: %.2f%%\", $3*100/$2 }'"),
		Disk: executeCmd("bash", "-c", "df -h | awk '$NF==\"/\"{printf \"Disk: %d/%dGB %s\", $3,$2,$5}'"),
	}
}

func main() {
	// Initialize the OLED with specific settings
	oled, err := goi2coled.NewI2c(ssd1306.SSD1306_SWITCHCAPVCC, 64, 128, 0x3C, 1)
	if err != nil {
		panic(err)
	}
	defer oled.Close()

	black := color.RGBA{0, 0, 0, 255}
	colWhite := color.RGBA{255, 255, 255, 255}

	drawer := &font.Drawer{
		Dst:  oled.Img,
		Src:  &image.Uniform{colWhite},
		Face: basicfont.Face7x13,
	}

	done := make(chan os.Signal, 1)
	stopCh := make(chan bool, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
				// Get stats
				stat := GetStat()
				
				lineSpacing := 9 * 32
				// Clear only the area where the stats are drawn
				draw.Draw(oled.Img, image.Rect(0, 16, 128, 64), &image.Uniform{black}, image.Point{}, draw.Src)
				
				// Display Stats
				drawer.Dot = fixed.Point26_6{X: 0, Y: fixed.Int26_6(10 * 64)}
				drawer.DrawString(fmt.Sprintf("IP: %s", stat.Ip))

				drawer.Dot = fixed.Point26_6{X: 0, Y: fixed.Int26_6(20*64 + lineSpacing)}
				drawer.DrawString(stat.Cpu)

				drawer.Dot = fixed.Point26_6{X: 0, Y: fixed.Int26_6(30*64 + 2 * lineSpacing)}
				drawer.DrawString(stat.Mem)

				drawer.Dot = fixed.Point26_6{X: 0, Y: fixed.Int26_6(40*64 + 3 * lineSpacing)}
				drawer.DrawString(stat.Disk)

				// Update display without flickering
				oled.Draw()
				oled.Display()

				time.Sleep(5 * time.Second)
			}
		}
	}()

	<-done
	stopCh <- true
	fmt.Println("Stop programme")
}
