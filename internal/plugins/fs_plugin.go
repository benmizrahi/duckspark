package plugins

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/benmizrahi/duckspark/internal/common"
	"github.com/benmizrahi/duckspark/internal/protos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FSPlugin struct {
	Path        string
	Format      string
	Parallelism int
}

// Execute implements contract.IPluginContract (worker job)
func (FSPlugin) Execute(task *protos.Task) *protos.TaskResult {

	from := task.Instactions[0]
	d, err := os.Open(from)
	if err != nil {
		logrus.Error("unable to read partition file", err)
		return &protos.TaskResult{
			Uuid:    task.Uuid,
			Status:  false,
			EndTime: timestamppb.Now(),
		}
	}
	defer d.Close()

	csvReader := csv.NewReader(d)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	rows := []*protos.DataRow{}
	for _, row := range data {

		d := protos.DataRow{
			Data: []string{},
		}
		for _, column := range row {
			d.Data = append(d.Data, column)
		}
		rows = append(rows, &d)
	}

	return &protos.TaskResult{
		Uuid:     task.Uuid,
		Status:   true,
		Dataflow: false,
		Data:     rows,
		EndTime:  timestamppb.Now(),
	}
}

// Configs implements contract.IPluginContract
func (p FSPlugin) Configs(conf map[string]string) common.IPluginContract {
	p.Format = conf["format"]
	p.Path = conf["path"]
	p.Parallelism = 1
	marks, err := strconv.ParseInt(conf["parallelism"], 10, 0)
	if err == nil {
		p.Parallelism = int(marks)
	}
	return p
}

// Plan implements plugins.IPluginContract
func (p FSPlugin) Plan(args ...interface{}) []*protos.Task {

	path := args[len(args)-1].(string)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: improve distribution
	distribution := []*protos.Task{}
	for _, file := range files {
		distribution = append(distribution, &protos.Task{
			Uuid:         uuid.New().String(),
			Instactions:  []string{path + file.Name()},
			Plugin:       p.Name(),
			CreationTime: timestamppb.Now(),
		})
	}

	return distribution
}

// Name implements plugins.IPluginContract
func (p FSPlugin) Name() string {
	return "fsplugin"
}

// Name must be New + struct name
func NewFSPlugin() common.IPluginContract {
	return FSPlugin{}
}
