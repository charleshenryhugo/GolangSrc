package consts

type (
	IZ  int
	FZ  float64
	STR string
	ByteSize float64
	char byte
)

const (
	Ln2                        = 0.693147180559945309417232121458176568075500134360255254120680009
	Pi                         = 3.1415926
	Log2E                      = 1 / Ln2
	Billion                    = 1e9
	HardEight                  = (1 << 100) >> 97
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

const (
	i1 = iota
	RED
	ORANGE
	YELLOW
	GREEN
	BLUE
	INDIGO
	VIOLET
)

const (
	_ = iota
	KB ByteSize = 1 << (10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
