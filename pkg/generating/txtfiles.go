package generating

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"path"
//	"strconv"
//	"sync"
//	"worder/custom"
//	"worder/utils"
//)
//
//const txtExtension = "txt"
//
//type TxtFilesGenerator struct {
//	src   string
//	dest  string
//	count int
//	size  int
//}
//
//func NewTxtFilesGenerator(
//	src, dest string, count, size int) *TxtFilesGenerator {
//	return &TxtFilesGenerator{
//		src:   src,
//		dest:  dest,
//		count: count,
//		size:  size,
//	}
//}
//
//func (tfg *TxtFilesGenerator) Generate() error {
//	wg := sync.WaitGroup{}
//	wg.Add(tfg.count)
//
//	source, err := os.Open(tfg.src)
//	defer source.Close()
//	if err != nil {
//		return custom.SampleFileLoadingErr
//	}
//
//	reader := bufio.NewReader(source)
//	sample, err := utils.GetFileContent(reader)
//	if err != nil {
//		return custom.SampleFileLoadingErr
//	}
//
//	for i := 0; i < tfg.count; i++ {
//		targetPath := path.Join(tfg.dest, fmt.Sprintf("%s.%s", strconv.Itoa(i), txtExtension))
//		target, err := os.Create(targetPath)
//		if err != nil {
//			fmt.Println(err)
//		}
//		go func(wg *sync.WaitGroup, file *os.File, payload string) {
//			err := tfg.write(file, payload)
//			if err != nil {
//				fmt.Println(err)
//			}
//			wg.Done()
//		}(&wg, target, sample)
//	}
//
//	wg.Wait()
//	return nil
//}
//
//func (tfg *TxtFilesGenerator) writeT(file *os.File, payload string) error {
//	defer file.Close()
//	writer := bufio.NewWriter(file)
//
//	for curSize := 0; curSize < tfg.size; {
//		written, err := writer.WriteString(payload)
//		if err != nil {
//			return err
//		}
//		curSize += written
//	}
//	err := writer.Flush()
//	if err != nil {
//		return err
//	}
//	return nil
//}
