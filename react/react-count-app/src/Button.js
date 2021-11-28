import React,{useState} from 'react';


function Button(props) {

    const [count,setCount]= useState(0);

    return (
        <div className="counter">
            <h3>{props.title_text}</h3>
            <p>カウント：{count}</p>
            <button onClick={()=>{setCount(count+1)}}>{props.add_text}</button>
            <button onClick={()=>{setCount(count-1)}}>{props.sub_text}</button>
        </div>
    )
}

export default Button
