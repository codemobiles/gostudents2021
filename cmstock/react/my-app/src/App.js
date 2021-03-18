import React from 'react'
import Axios from 'axios'

export default function App() {

  const [youtubes, setYoutubes] = React.useState([])

  React.useEffect(async () => {
    const url = "https://codemobiles.com/adhoc/youtubes/index_new.php?username=admin&password=password&type=foods"
    const result = await Axios.get(url)
    console.log(result.data.youtubes[0].title)
    setYoutubes(result.data.youtubes)
  }, [])

  return (
    <div>
      <h1>CodeMobiles</h1>
      <ul>
        {youtubes.map(item=><li key={item.id}>{item.title}</li>)}         
      </ul>
    </div>
  )
}
