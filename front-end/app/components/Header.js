import Link from 'next/link'
import React from 'react'

const Header = () => {
  return (
    <div className=' top-0 w-full h-12 bg-[#1e1e1e] border-b border-[#6b6b6b] px-4'>
      <ul className='flex items-center h-full gap-[30px] text-white text-sm'>
        <Link href="/" className='hover:underline hover:underline-offset-4'>Home</Link>
        <Link href="/route" className='hover:underline hover:underline-offset-4'>Create/Update Route </Link>
        <Link href="/route-create" className='hover:underline hover:underline-offset-4'>Get/Delete Route</Link>
        <Link href="/type-data" className='hover:underline hover:underline-offset-4'>Create/Update Type</Link>
        <Link href="/flow" className='hover:underline hover:underline-offset-4'>Flow</Link>
        <Link href="/upload-csv" className='hover:underline hover:underline-offset-4'>Upload csv</Link>
      </ul>
    </div>
  )
}

export default Header