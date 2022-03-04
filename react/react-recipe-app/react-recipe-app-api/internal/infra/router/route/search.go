package route

import (
	"kedamaonsen/react-recipe-app-api/internal/domain/entity"
	"kedamaonsen/react-recipe-app-api/internal/infra/database"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func Search(c echo.Context) error {

	//クエリ取得
	query := c.QueryParam("q")

	//""で囲まれてたらそれを削除
	query = strings.ReplaceAll(query, `"`, "")

	//検索結果
	var recipes []entity.Recipe

	//レシピ集から、タイトルに検索文字が含まれるものだけ取得
	for _, recipe := range database.Database.Recipes {

		if strings.Contains(recipe.Title, query) {
			recipes = append(recipes, recipe)
		}
	}

	result := map[string]interface{}{}
	result["hits"] = recipes

	return c.JSON(http.StatusOK, result)
}
