package main

/*
import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

//WriteDflt Write a Defaults struct into a target file
func (dflt *Defaults) WriteDflt(filename string) error {
	//Write the struct into a []byte
	output, err := xml.MarshalIndent(&dflt, "", "  ")
	if err != nil {
		return err
	}
	//Open the to-be-written file
	xmlFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	//Write the target file
	if err := ioutil.WriteFile(filename, output, 0777); err != nil {
		return err
	}

	//If everything worked well, return nil
	return nil
}

//SetDfltMsg updates default message
func (dflt *Defaults) SetDfltMsg(newMsg, filename string) error {
	//Update default message
	dflt.Message = newMsg
	return dflt.WriteDflt(filename)
}

func (dflt *Defaults) SetDfltSbjt(newSbjt, filename string) error {
	//Update default subject
	dflt.Subject = newSbjt
	return dflt.WriteDflt(filename)
}

func (dflt *Defaults) SetDfltMsgFile(newMsgFile, filename string) error {
	//Update default message file
	dflt.MessageFile = newMsgFile
	return dflt.WriteDflt(filename)
}

func (dflt *Defaults) SetDfltEmailListFile(newListFile, filename string) error {
	dflt.EmailListFile = newListFile
	return dflt.WriteDflt(filename)
}

func (dflt *Defaults) SetDfltSlackListFile(newListFile, filename string) error {
	dflt.SlackListFile = newListFile
	return dflt.WriteDflt(filename)
}
*/
