import React from 'react'

const Recipe = (props) => {
    return (
        <div>
            <h1>{props.title}</h1>
            <ol>
                {props.ingredients.map(ingredients => (
                    <li key={ingredients}>
                        {ingredients}
                    </li>
                ))}
            </ol>
            <p>{props.calories} cal</p>
            <img src="" className="dummy" alt=""></img>
        </div>
    )
}

export default Recipe
