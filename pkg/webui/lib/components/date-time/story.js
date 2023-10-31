

/* eslint-disable react/prop-types */

import React from 'react'

import RelativeDateTime from './relative'

const Example = ({ title, unit, ago }) => {
  const date = new Date()
  if (unit === 'hour') {
    date.setHours(date.getHours() - ago)
  } else if (unit === 'minute') {
    date.setMinutes(date.getMinutes() - ago)
  } else if (unit === 'day') {
    date.setDate(date.getDate() - ago)
  } else if (unit === 'month') {
    date.setMonth(date.getMonth() - ago)
  } else if (unit === 'year') {
    date.setFullYear(date.getFullYear() - ago)
  }

  return (
    <div>
      <h3>{title}:</h3>
      <RelativeDateTime value={date} />
      <hr />
    </div>
  )
}

export default {
  title: 'DateTime/Relative',
}

export const Default = () => (
  <div>
    <Example title="from now" />
    <Example title="from 1 minute ago" unit="minute" ago={1} />
    <Example title="from 30 minutes ago" unit="minute" ago={30} />
    <Example title="from 1 hour ago" unit="hour" ago={1} />
    <Example title="from 12 hours ago" unit="hour" ago={12} />
    <Example title="from 1 day ago" unit="day" ago={1} />
    <Example title="from 7 days ago" unit="day" ago={7} />
    <Example title="from 1 month ago" unit="month" ago={1} />
    <Example title="from 6 months ago" unit="month" ago={6} />
    <Example title="from 1 year ago" unit="year" ago={1} />
    <Example title="from 2 years ago" unit="year" ago={2} />
  </div>
)
