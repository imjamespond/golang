import React from 'react'

import logo from '@assets/logo192.png'
import './TestChunk.css'

export default function(){

  return <div>
    test  chunk
    <div className='logo' style={{width:192,height:192}}></div>
    <img src={logo} />
  </div>
}