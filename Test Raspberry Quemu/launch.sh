launch_vm1="
qemu-system-aarch64 \
    -M raspi3b \
    -cpu cortex-a72 \
    -append \"rw earlyprintk loglevel=8 console=ttyAMA0,115200 dwc_otg.lpm_enable=0 root=/dev/mmcblk0p2 rootdelay=1\" \
    -dtb bcm2710-rpi-3-b-plus.dtb \
    -sd disk_1.img \
    -kernel kernel8.img \
    -m 1G -smp 4 \
    -nographic \
    -usb -device usb-mouse -device usb-kbd \
    -device usb-net,netdev=net0 \
    -netdev user,id=net0,hostfwd=tcp::5555-:22 \
"

launch_vm2="
qemu-system-aarch64 \
    -M raspi3b \
    -cpu cortex-a72 \
    -append \"rw earlyprintk loglevel=8 console=ttyAMA0,115200 dwc_otg.lpm_enable=0 root=/dev/mmcblk0p2 rootdelay=1\" \
    -dtb bcm2710-rpi-3-b-plus.dtb \
    -sd disk_2.img \
    -kernel kernel8.img \
    -m 1G -smp 4 \
    -nographic \
    -usb -device usb-mouse -device usb-kbd \
    -device usb-net,netdev=net0 \
    -netdev user,id=net0,hostfwd=tcp::5556-:22 \
"

# Open a new terminal window and execute the first vm
gnome-terminal -- bash -c "$launch_vm1; exec bash"

# Open a new terminal window and execute the second vm
gnome-terminal -- bash -c "$launch_vm2; exec bash"



