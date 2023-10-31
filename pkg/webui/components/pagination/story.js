

import React from 'react'
import { withInfo } from '@storybook/addon-info'

import Pagination from '.'

export default {
  title: 'Pagination',

  decorators: [
    withInfo({
      inline: true,
      header: false,
      propTables: [Pagination],
    }),
  ],
}

export const Default = () => (
  <div>
    <Pagination pageCount={1} />
    <Pagination pageCount={3} initialPage={2} marginPagesDisplayed={2} />
    <Pagination pageCount={3} initialPage={3} marginPagesDisplayed={2} />
  </div>
)

export const AllPagesWithoutGaps = () => <Pagination pageCount={10} pageRangeDisplayed={10} />

AllPagesWithoutGaps.story = {
  name: 'All pages (without gaps)',
}

export const WithGaps = () => (
  <div>
    <Pagination pageCount={20} marginPagesDisplayed={2} />
    <Pagination pageCount={9} initialPage={4} pageRangeDisplayed={1} marginPagesDisplayed={2} />
    <Pagination pageCount={9} initialPage={4} pageRangeDisplayed={3} marginPagesDisplayed={2} />
  </div>
)

WithGaps.story = {
  name: 'With gaps',
}
