import 'core-js/stable'
import React from 'react'
import ReactDom from 'react-dom'

import App from './App'

interface Person {
  name: string
  age: number
  gender?: string
}

const someone: Person = {
  name: 'foobar',
  age: 99,
  gender: 'unknown'
}
console.log(someone)

ReactDom.render(<App/>, document.getElementById('app'))
