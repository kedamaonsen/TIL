package database

import (
	"io/ioutil"
	"kedamaonsen/react-recipe-app-api/internal/domain/entity"

	"gopkg.in/yaml.v3"
)

type database struct {
	Recipes []entity.Recipe
}

var Database database

func init() {
	buf, err := ioutil.ReadFile("./data.yml")
	if err != nil {
		panic(err)
	}

	//config全体取得
	var conf map[string]interface{}
	yaml.Unmarshal(buf, &conf)

	//recipes配下取得
	recipes := conf["recipes"]

	//配下を再度バイト変換
	b, _ := yaml.Marshal(recipes)

	//レシピ集取得
	yaml.Unmarshal(b, &Database.Recipes)
}
