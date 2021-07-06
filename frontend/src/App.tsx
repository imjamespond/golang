import React from 'react'

import './App.css'

export default function(){

  const [chunk, setChunck] = React.useState<any>()

  return <div>
    hello react...
    <br/>
    <button onClick={()=>{
      import('./test/TestChunk')
      .then(({ default: Chunk }) => {
        Chunk && setChunck(<Chunk/>)
      })
    }}>test chunk</button>
    {chunk}
  </div>
}