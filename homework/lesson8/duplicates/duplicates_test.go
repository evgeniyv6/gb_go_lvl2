package duplicates

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	testFolders    = "./testfolder/inner/"
	file1      = "./testfolder/f1"
	file2      = "./testfolder/f2"
	innerFile = "./testfolder/inner/f3"
	NEWFS  = afero.NewOsFs() // from afero ref.: In the case of OsFs it
							// will still use the same underlying filesystem
							// but will reduce the ability to drop in other filesystems as desired.
	AFS  = &afero.Afero{Fs: NEWFS}
)

func init() {
	// create test files and directories
	AFS.MkdirAll(testFolders, 0755)
	afero.WriteFile(AFS, file1, []byte("test file"), 0755)
	afero.WriteFile(AFS, file2, []byte("test file"), 0755)
	afero.WriteFile(AFS, innerFile, []byte("test file"), 0755)
}

var fsTest fileSystem = osFSMock{}
type osFSMock struct{}
func (osFSMock) Open(name string) (ifile, error) {
	return AFS.Open(name)
}
func (osFSMock) Stat(name string) (os.FileInfo, error) {
	return AFS.Stat(name)
}

// go test -v .
func TestFindDuplicates(t *testing.T) {
	test := struct {
		in string
		want map[string][]string
	}{
		"./testfolder/",
		map[string][]string{"0000000000000009:F2646BC1":
			{"testfolder/f2","testfolder/inner/f3","testfolder/f1"}},
	}
	ch := make(chan FileStat, 10)

	go FindDuplicates(ch,test.in, fsTest)
	data := MapResults(ch)
	_ = ResultWorker(data, false)

	mm := make(map[string][]string, len(data))
	for k,v:=range data {
		mm[k] = *v
	}

	for k, _ := range test.want {
		assert.ElementsMatch(t, test.want[k], mm[k])
	}
}