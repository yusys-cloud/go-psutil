// Author: yangzq80@gmail.com
// Date: 2021-11-29
//
package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(memCmd)
	rootCmd.AddCommand(netCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(diskCmd)
	rootCmd.AddCommand(cpuCmd)
	rootCmd.AddCommand(netInterfaceCmd)
}

var memCmd = &cobra.Command{
	Use:   "mem",
	Long:  `mem info`,
	Short: "内存使用",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")

		v, _ := mem.VirtualMemory()

		fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total/1024/1024/1024/1024, v.Free/1024/1024/1024, v.UsedPercent)

		fmt.Println(v)
	},
}

var netCmd = &cobra.Command{
	Use:   "net",
	Long:  `net info`,
	Short: "网卡列表",
	Run: func(cmd *cobra.Command, args []string) {

		list, _ := net.Interfaces()

		for _, o := range list {
			log.Println(o)
		}

		log.Println(net.IOCountersStat{})
	},
}

var netInterfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "网卡名称",
	Run: func(cmd *cobra.Command, args []string) {

		list, _ := net.Interfaces()

		for _, o := range list {
			if o.HardwareAddr != "" && len(o.Flags) > 0 && o.Flags[0] == "up" && len(o.Addrs) > 0 {
				for _, addr := range o.Addrs {
					if strings.Contains(addr.Addr, ".") {
						ip := strings.Split(addr.Addr, "/")[0]
						//log.Println(ip, addr)
						lastIp := ip[strings.LastIndex(ip, ".")+1 : len(ip)]
						lastIpInt, _ := strconv.Atoi(lastIp)
						if lastIpInt > 1 {
							fmt.Println(o.Name)
						}
					}
				}
			}
		}
	},
}

var loadCmd = &cobra.Command{
	Use:   "load",
	Long:  `load info`,
	Short: "CPU负载",
	Run: func(cmd *cobra.Command, args []string) {

		stat, _ := load.Avg()

		fmt.Println(stat.String())
	},
}

var diskCmd = &cobra.Command{
	Use:  "disk",
	Long: `disk info`,
	Run: func(cmd *cobra.Command, args []string) {

		stat, _ := disk.Usage("/")

		fmt.Println(stat)
	},
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Long:  `cpu info`,
	Short: "CPU信息",
	Run: func(cmd *cobra.Command, args []string) {

		stat, _ := cpu.Info()

		fmt.Println(stat)
	},
}
