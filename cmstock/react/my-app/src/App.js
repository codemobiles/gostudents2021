import React from 'react'

export default function App() {
  return (
    <div>
      <h1>CodeMobiles</h1>
      <ul>
        {[1,2,3,4].map(item=><li key={item.toString()}>{item}</li>)}         
      </ul>
    </div>
  )
}
