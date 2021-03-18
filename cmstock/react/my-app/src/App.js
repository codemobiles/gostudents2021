import React from 'react'
import Axios from 'axios'

export default function App() {

  const [youtubes, setYoutubes] = React.useState([])
  const [products, setProducts] = React.useState([])

  /*
  React.useEffect(async () => {
    // const url = "https://codemobiles.com/adhoc/youtubes/index_new.php?username=admin&password=password&type=foods"
    const result = await Axios.get(url)
    console.log(result.data.youtubes[0].title)
    setYoutubes(result.data.youtubes)
  }, [])
  */

  React.useEffect(async () => {
    const url = "http://localhost:8081/api/v2/product"
    const result = await Axios.get(url)
    setProducts(result.data)
  }, [])




  return (
    <div>
      <h1>CodeMobiles</h1>
      <ul>
        {products.map(item=><li key={item.id}>{item.name}</li>)}         
      </ul>
    </div>
  )
}
