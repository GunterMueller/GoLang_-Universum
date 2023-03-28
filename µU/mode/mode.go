package mode

// (c) Christian Maurer   v. 210526 - license see µU.go

var (
  x, y [NModes]uint
  m Mode
  wn, hn = uint(0), uint(0)
)

func init() { //       framebuffer colourdepth: 8 bit 16 bit 24 bit
  m = None;    x[m], y[m] =    0,    0
  m = Mini;    x[m], y[m] =  192,  160 //  6:5
  m = HQVGA;   x[m], y[m] =  240,  160 //  4:3
  m = QVGA;    x[m], y[m] =  320,  240 //  4:3  0x334  0x335  0x336
  m = HVGA;    x[m], y[m] =  480,  320 //  4:3
  m = TXT;     x[m], y[m] =  640,  400 //  8:5  0x300  0x33d  0x33e (0x300 = 4 bit. no graphics)
  m = VGA;     x[m], y[m] =  640,  480 //  4:3  0x301  0x311  0x312
  m = PAL;     x[m], y[m] =  768,  576 //  4:3
  m = WVGA;    x[m], y[m] =  800,  480 //  5:3
  m = SVGA;    x[m], y[m] =  800,  600 //  4:3  0x303  0x314  0x315
  m = XGA;     x[m], y[m] = 1024,  768 //  4:3  0x305  0x317  0x318
  m = HD;      x[m], y[m] = 1280,  720 // 16:9
  m = WXGA;    x[m], y[m] = 1280,  800 //  8:5  0x366  0x367  0x368
  m = SXVGA;   x[m], y[m] = 1280,  960 //  4:3
  m = SXGA;    x[m], y[m] = 1280, 1024 //  5:4  0x307  0x31a  0x31b
  m = WXGA1;   x[m], y[m] = 1366,  768 // 16:9 ca.
  m = SXGAp;   x[m], y[m] = 1400, 1050 //  5:4  0x347  0x348  0x349 or 0x343  0x345  0x346 ?
  m = WXGAp;   x[m], y[m] = 1440,  900 //  8:5  0x364         0x365 or 0x369  0x36a  0x36b ?
  m = WXGApp;  x[m], y[m] = 1600,  900 // 16:9
  m = WSXGA;   x[m], y[m] = 1600, 1024 // 25:16
  m = UXGA;    x[m], y[m] = 1600, 1200 //  4:3  0x345  0x346  0x34a or 0x33a  0x34b  0x35a ?
  m = WSXGAp;  x[m], y[m] = 1680, 1050 // 25:16 0x368         0x369
  m = FHD;     x[m], y[m] = 1920, 1080 // 16:9
  m = WUXGA;   x[m], y[m] = 1920, 1200 //  8:5  0x37c         0x37d
  m = SUXGA;   x[m], y[m] = 1920, 1440 //  4:3  0x33c  0x34d  0x35c
  m = QWXGA;   x[m], y[m] = 2048, 1152 // 16:9
  m = QXGA;    x[m], y[m] = 2048, 1536 //  4:3                0x352
  m = WSUXGA;  x[m], y[m] = 2560, 1440 // 16:9
  m = WQXGA;   x[m], y[m] = 2560, 1600 //  8:5
  m = QSXGAp;  x[m], y[m] = 2800, 2100 //  4:3
  m = QUXGA;   x[m], y[m] = 3200, 2400 //  4:3
  m = UHD;     x[m], y[m] = 3840, 2160 // 16:9
  m = HXGA;    x[m], y[m] = 4096, 3072 //  4:3
  m = WHXGA;   x[m], y[m] = 5120, 3200 //  8:5
  m = HSXGA;   x[m], y[m] = 5120, 4096 //  5:4
  m = HUXGA;   x[m], y[m] = 6400, 4800 //  4:3
  m = FUHD;    x[m], y[m] = 7680, 4320 // 16:9
  m = NEW;     x[m], y[m] =   wn,   hn
}

func modeOf (w, h uint) Mode {
  for m := Mode(0); m + 1 < NModes; m++ {
    if x[m] == w && y[m] == h {
      return m
    }
  }
  m = NEW
  wn, hn = w, h
  x[m], y[m] = wn, hn
  return m
}
