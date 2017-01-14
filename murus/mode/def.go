package mode

// (c) murus.org  v. 160123 - license see murus.go

type
  Mode byte; const (
  UVGA = Mode(iota)
         //   5 x  15 /  120 x   80
  Mini   //  10 x  24 /  192 x  160
  MVGA   //  10 x  30 /  240 x  160
  QVGA   //  15 x  40 /  320 x  240
  HVGA   //  20 x  60 /  480 x  320
  TXT    //  25 x  80 /  640 x  400
  VGA    //  30 x  80 /  640 x  480
  PAL    //  36 x  96 /  768 x  576
  WVGA   //  30 x 100 /  800 x  480
  SVGA   //  37 x 100 /  800 x  600
  WPAL   //  37 x 120 /  960 x  600
  WVGApp //  36 x 128 / 1024 x  576
  WSVGA  //  37 x 128 / 1024 x  600
  XGA    //  48 x 128 / 1024 x  768
  HD     //  45 x 160 / 1280 x  720
  WXGA   //  50 x 160 / 1280 x  800
  SXVGA  //  60 x 160 / 1280 x  960
  SXGA   //  64 x 160 / 1280 x 1024
  WXGA1  //  48 x 171 / 1366 x  768
  SXGAp  //  65 x 175 / 1400 x 1050
  WXGAp  //  56 x 180 / 1440 x  900
  WXGApp //  56 x 200 / 1600 x  900
  WSXGA  //  64 x 200 / 1600 x 1024
  UXGA   //  75 x 200 / 1600 x 1200
  WSXGAp //  65 x 210 / 1680 x 1050
  FHD    //  67 x 240 / 1920 x 1080
  WUXGA  //  75 x 240 / 1920 x 1200
  SUXGA  //  90 x 240 / 1920 x 1440
  QWXGA  //  72 x 256 / 2048 x 1152
  QXGA   //  96 x 256 / 2048 x 1536
// ?     //  67 x 320 / 2560 x 1080
  QHD    //  90 x 320 / 2560 x 1440
  WQXGA  // 100 x 320 / 2560 x 1600
  QSXGA  // 128 x 320 / 2560 x 2048
  QSXGAp // 131 x 350 / 2800 x 2100
  QWXGApp// 112 x 400 / 3200 x 1800
  WQSXGA // 128 x 400 / 3200 x 2048
  QUXGA  // 150 x 400 / 3200 x 2400
  WQHD   //  90 x 430 / 3440 x 1440
  UHD    // 135 x 440 / 3840 x 2160
  WQUXGA // 150 x 480 / 3840 x 2400
  WUHD   // 135 x 512 / 4096 x 2160
  HXGA   // 192 x 512 / 4096 x 3072
  QQHD   // 188 x 640 / 5120 x 2880
  WHXGA  // 200 x 640 / 5120 x 3200
  HSXGA  // 256 x 640 / 5120 x 4096
  WHSXGA // 256 x 800 / 6400 x 4096
  HUXGA  // 300 x 800 / 6400 x 4800
  FUHD   // 270 x 960 / 7680 x 4320
  WHUXGA // 300 x 960 / 7680 x 4800
  NModes )

func Wd (m Mode) uint { return nX[m] }

func Ht (m Mode) uint { return nY[m] }

func Res (m Mode) (uint, uint) { return nX[m], nY[m] }

func ModeOf (w, h uint) Mode { return modeOf(w,h) }
