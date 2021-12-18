import React, { useEffect, useState ,useRef } from 'react'
import './App.css';
import Recipe from './Recipe';

function App() {

  const inputRef = useRef(null);
  useEffect(() => {
    inputRef.current.focus();
  },[])

  const [search, setSearch] = useState("");
  const [query, setQuery] = useState("");
  const [recipes, setRecipes] = useState([]);

  const getRecipes = async () => {
    const response = await fetch(`http://localhost:9999/search?q="${query}"`);
    const data = await response.json();
    setRecipes(data.hits);
  }

  const getSearch = e =>{

    //reactで検索処理したいので、htmlのformクリックイベントはキャンセルする
    e.preventDefault();

    //queryに値設定
    setQuery(search);

    //検索窓の値クリア
    setSearch('');
  }

  const updateSearch = e => {
    setSearch(e.target.value);
  }

  //queryが変更されたらレシピ検索
  useEffect(() => {
    getRecipes();
  }, [query])

  return (
    <div className="App">
      <form onSubmit={getSearch}>
        <input ref={inputRef} type="text" value={search} onChange={updateSearch}></input>
        <button type="submit">検索</button>
      </form>
      <div>
        {recipes.map(recipe => (
          <Recipe
          key={recipe.Title}
          title={recipe.Title}
          calories={recipe.Calory}
          ingredients={recipe.Ingredients}
          />
        ))}
      </div>
    </div>
  );
}

export default App;
