package mode

// (c) murus.org  v. 160726 - license see murus.go

import (
  "strconv"
  "murus/ker"
)
var
  nX, nY [NModes]uint

func init() {
//  Lines x Columns          Settings in /boot/grub/menu.lst
//  at Font 16x8             at colourdepth
//                           8 bit:        16 bit:       24 bit:
//                           vga =         vga =         vga =

//      5 x  15       4:3
  nX[UVGA] =    120
  nY[UVGA] =     80

//     10 x  24       6:5
  nX[Mini] =    192
  nY[Mini] =    160

//     10 x  30       4:3
  nX[MVGA] =    240
  nY[MVGA] =    160
/*
//     12 x  40       8:5    0x330 (816)   0x30e (782)   0x30f (783)
  nX[    ] =    320
  nY[    ] =    200
*/
//     15 x  40       4:3    0x334 (820)   0x335 (821)   0x336 (822)
  nX[QVGA] =    320
  nY[QVGA] =    240
/*
//     25 x  40       4:5    0x331 (817)   0x332 (818)   0x333 (819)
  nX[    ] =    320
  nY[    ] =    400
*/
//     20 x  60       4:3
  nX[HVGA] =    480
  nY[HVGA] =    320

//                           4 bit: (no graphics)
//     25 x  80       8:5    0x300 (768)   0x33d (829)   0x33e (830)
  nX[TXT] =     640
  nY[TXT] =     400

//     30 x  80       4:3    0x301 (769)   0x311 (785)   0x312 (786)
  nX[VGA] =     640
  nY[VGA] =     480

//     36 x  96       4:3
  nX[PAL] =     768
  nY[PAL] =     576

//     30 x 100       5:3
  nX[WVGA] =    800
  nY[WVGA] =    480

//     37 x 100       4:3    0x303 (771)   0x314 (788)   0x315 (789)
  nX[SVGA] =    800
  nY[SVGA] =    600

//     37 x 120       8:5    0x363 (867)   0x364 (868)   0x365 (869)
  nX[WPAL] =    960
  nY[WPAL] =    600

//     36 x 128  etwa 5:3
  nX[WVGApp] = 1024
  nY[WVGApp] =  576

//     37 x 128  etwa 5:3
  nX[WSVGA] =  1024
  nY[WSVGA] =   600

//     48 x 128       4:3    0x305 (773)   0x317 (791)   0x318 (792)
  nX[XGA] =    1024
  nY[XGA] =     768

//     45 x 160      16:9
  nX[HD] =     1280
  nY[HD] =      720

//     50 x 160       8:5    0x366 (870)   0x367 (871)   0x368 (872)
  nX[WXGA] =   1280 // also 1366
  nY[WXGA] =    800 // also  768

//     60 x 160       8:5
  nX[SXVGA] =  1280
  nY[SXVGA] =   800

//     64 x 160       5:4    0x307 (775)   0x31a (794)   0x31b (795)

  nX[SXGA] =   1280
  nY[SXGA] =   1024

//     48 x 171      16:9 ca.
  nX[WXGA1] =  1366
  nY[WXGA1] =   768

//     65 x 175       5:4    0x347 (839)   0x348 (840)   0x349 (841)
//                     or ?  0x343 (835)   0x345 (837)   0x346 (838)
  nX[SXGAp] =  1400
  nY[SXGAp] =  1050

//     56 x 180       8:5    0x364 (868)                 0x365 (869)
//                     or ?  0x369 (873)   0x36a (874)   0x36b (875)
  nX[WXGAp] =  1440
  nY[WXGAp] =   900

//     56 x 200      16:9
  nX[WXGApp] = 1600
  nY[WXGApp] =  900

//     64 x 200      25:16
  nX[WSXGA] =  1600
  nY[WSXGA] =  1024

//     75 x 200       4:3    0x345 (837)   0x346 (838)   0x34a (842)
  nX[UXGA] =   1600
  nY[UXGA] =   1200

//     65 x 210      25:16   0x368 (872)                 0x369 (873)
  nX[WSXGAp] = 1680
  nY[WSXGAp] = 1050

//     67 x 240      16:9
  nX[FHD] =    1920
  nY[FHD] =    1080

//     75 x 240       8:5    0x37c (892)                 0x37d (893)
  nX[WUXGA] =  1920
  nY[WUXGA] =  1200

//     90 x 240       4:3
  nX[SUXGA] =  1920
  nY[SUXGA] =  1440

//     72 x 256      16:9
  nX[QWXGA] =  2048
  nY[QWXGA] =  1152

//     96 x 256       4:3                                0x352 (850)
  nX[QXGA] =   2048
  nY[QXGA] =   1536

//     67 x 320      64:27
//nX[ ?  ] =   2560
//nY[ ?  ] =   1080

//     90 x 320      16:9
  nX[QHD]  =   2560
  nY[QHD]  =   1440

//    100 x 320       8:5
  nX[WQXGA] =  2560
  nY[WQXGA] =  1600

//    128 x 320       5:4
  nX[QSXGA] =  2560
  nY[QSXGA] =  2048

//    131 x 350       4:3
  nX[QSXGAp] = 2800
  nY[QSXGAp] = 2100

//    112 x 400      16:9
  nX[QWXGApp] = 3200
  nY[QWXGApp] = 1800

//    128 x 400      25:16
  nX[WQSXGA] = 3200
  nY[WQSXGA] = 2048

//    150 x 400       4:3
  nX[QUXGA] =  3200
  nY[QUXGA] =  2400

//     90 x 430      43:18
  nX[WQHD] =   3440
  nY[WQHD] =   1440

//    135 x 440      16:9
  nX[UHD] =    3840
  nY[UHD] =    2160

//    150 x 480       8:5
  nX[WQUXGA] = 3840
  nY[WQUXGA] = 2400

//    135 x 512
  nX[WUHD] =   4096
  nY[WUHD] =   2160

//    192 x 512       4:3
  nX[HXGA] =   4096
  nY[HXGA] =   3072

//    188 x 640
  nX[QQHD]  =  5120
  nY[QQHD]  =  2880

//    200 x 640       8:5
  nX[WHXGA] =  5120
  nY[WHXGA] =  3200

//    256 x 640       5:4
  nX[HSXGA] =  5120
  nY[HSXGA] =  4096

//    256 x 800      25:16
  nX[WHSXGA] = 6400
  nY[WHSXGA] = 4096

//    300 x 800       4:3
  nX[HUXGA] =  6400
  nY[HUXGA] =  4800

//    270 x 960      16:9
  nX[FUHD] =   7680
  nY[FUHD] =   4320

//    300 x 960       8:5
  nX[WHUXGA] = 7680
  nY[WHUXGA] = 4800
}

func modeOf (w, h uint) Mode {
  for m:= Mode(0); m < NModes; m++ {
    if nX[m] == w && nY[m] == h {
      return m
    }
  }
  ker.Panic ("hardware reports undefined mode " + strconv.Itoa(int(w)) + " x " + strconv.Itoa(int(h)))
  return NModes
}
