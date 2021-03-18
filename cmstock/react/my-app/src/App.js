import React from 'react'
import Axios from 'axios'

export default function App() {


  React.useEffect(async () => {
    const url = "https://codemobiles.com/adhoc/youtubes/index_new.php?username=admin&password=password&type=foods"
    const result = await Axios.get(url)
    console.log(result.data)
  }, [])

  return (
    <div>
      <h1>CodeMobiles</h1>
      <ul>
        {[1,2,3,4].map(item=><li key={item.toString()}>{item}</li>)}         
      </ul>
    </div>
  )
}
