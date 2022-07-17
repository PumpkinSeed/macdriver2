package main

import (
	"fmt"
	"github.com/PumpkinSeed/macdriver2/corebluetooth"
	"github.com/PumpkinSeed/macdriver2/objc"
)

func main() {
	centralManagerDelegate := objc.NewClass("CBCentralManagerDelegate", "NSObject")
	centralManagerDelegate.AddMethod("centralManager:didDiscoverPeripheral:advertisementData:RSSI:", func(
		central objc.Object,
		peripheral objc.Object,
		advertisementData objc.Object,
		rssi objc.Object,
	) {
		fmt.Println(central)
	})
	objc.RegisterClass(centralManagerDelegate)

	centralManagerDelegateInstance := objc.Get("CBCentralManagerDelegate").Alloc().Init()

	centralManager := corebluetooth.CBCentralManager{}
	centralManager.SetDelegate_(centralManagerDelegateInstance)

	centralManager.ScanForPeripheralsWithServices_options_(nil, nil)
}
