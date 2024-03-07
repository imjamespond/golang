type Entry<T> = { [K in keyof T]: [K, T[K]] }[keyof T];
export function objectEntries<T extends object>(t: T) {
  return Object.entries(t) as Entry<T>[]; // you need some assertion here
  // if TS allowed this without "as any" you wouldn't need this function at all
}

// const obj: Record<'foo'|'bar', number> = {
//   foo: 0,
//   bar: 0
// }
// objectEntries(obj).forEach(([k,v])=>{

// })