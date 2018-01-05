package term

// (c) Christian Maurer   v. 171229 - license see nU.go

// #include <termios.h>
// #include <unistd.h>
// #include <fcntl.h>
// #include <sys/ioctl.h>
// #include <linux/kd.h>
/*
struct termios t;
static int k;

void new (void) {
  ioctl (0, TCGETS, &t);
  tcgetattr (0, &t);
  ioctl (0, KDGKBMODE, &k);
  struct termios t1;
  t1 = t;
  t1.c_iflag = 0;
  t1.c_lflag &= ~(ISIG | ICANON | ECHO );
  t1.c_cc[VTIME] = 1;
  t1.c_cc[VMIN] = 3;
  tcsetattr (0, TCSANOW, &t1);
  ioctl (0, KDSKBMODE, K_MEDIUMRAW);
}

void tfin (void) {
  tcsetattr (0, TCSAFLUSH, &t);
  ioctl (0, KDSKBMODE, k);
}

char tread (void) {
  char b;
  read (0, &b, 1);
  return b;
}
*/
import "C"

type terminal struct {
  bool "active"
  byte
}

var a *terminal

func (x *terminal) Read() byte {
  x.byte = byte(C.tread())
  if x.byte == Esc {
    x.Fin()
  }
  return x.byte
}

func new_() Terminal {
  x := new(terminal)
  C.new()
  x.bool = true
  a = x
  return x
}

func (x *terminal) Fin() {
  if x.bool {
    C.tfin()
    x.bool = false
  }
}
