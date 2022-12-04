package data

// 周波数帯
type BandFreq string

const (
	Band19   BandFreq = "1.9"
	Band35   BandFreq = "3.5"
	Band7    BandFreq = "7"
	Band10   BandFreq = "10"
	Band14   BandFreq = "14"
	Band18   BandFreq = "18"
	Band21   BandFreq = "21"
	Band24   BandFreq = "24"
	Band28   BandFreq = "28"
	Band50   BandFreq = "50"
	Band144  BandFreq = "144"
	Band430  BandFreq = "430"
	Band1200 BandFreq = "1200"
	Band2400 BandFreq = "2400"
	Band5600 BandFreq = "5600"
	Band10G  BandFreq = "10G"
)

type Band struct {
	Freq   BandFreq
	Points PointElement
}

// SSB, CW, FM, AMの各得点。インデックスはdata.CommModeに同じ
type PointElement [4]int

type ContestConfig struct {
	Bands map[BandFreq]Band
}
