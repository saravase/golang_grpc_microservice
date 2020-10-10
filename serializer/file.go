package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WriteProtoBufToJSONFile(message proto.Message, filename string) error {

	data, err := ProtoBufToJSON(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto buf message into json file : %v", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("Cannot write json data to file : %v", err)
	}

	return nil
}

func WriteProtoBufToBinaryFile(message proto.Message, filename string) error {

	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto buf message into binary file : %v", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("Cannot write binary data to file : %v", err)
	}

	return nil
}

func ReadProtoBufFromBinaryFile(filename string, message proto.Message) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read binary data from file : %v", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Cannot Unmarshal proto buf message from binary file : %v", err)
	}

	return nil
}
