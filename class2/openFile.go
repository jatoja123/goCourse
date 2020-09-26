package class2
import ( "os"; "log"; "bufio" )

//	funkcja		argumenty	   zwraca
func OpenFile(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalf("%v",err)
	}

	result := make([]byte,1000)
	file.Read(result)

	return string(result)
}

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatalf("%v",err)
	}

	scammer := bufio.NewScanner(file)

	result := []string{}
	for scammer.Scan(){
		result = append(result, scammer.Text())
	}

	return result
}