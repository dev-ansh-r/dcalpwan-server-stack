

import React from 'react'

import Link from '.'

const divStyle = { width: '100px', height: '100px', backgroundColor: 'lightblue' }

const titleStyle = { marginRight: '20px' }

export default {
  title: 'Link',
}

export const DefaultUnstyled = () => (
  <div>
    <div>
      <span style={titleStyle}>link:</span>
      <Link to="/">Show more</Link>
    </div>
    <div>
      <span style={titleStyle}>anchor link:</span>
      <Link.Anchor href="/">Show more</Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>external anchor link:</span>
      <Link.Anchor external href="/">
        More information
      </Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>doc link:</span>
      <Link.DocLink path="/getting-started">Getting Started</Link.DocLink>
    </div>
  </div>
)

DefaultUnstyled.story = {
  name: 'Default (Unstyled)',
}

export const AsWrapper = () => (
  <div>
    <div>
      <span style={titleStyle}>linked block element:</span>
      <Link to="/">
        <div style={divStyle} />
      </Link>
    </div>
    <div>
      <span style={titleStyle}>anchor linked block element:</span>
      <Link.Anchor href="/">
        <div style={divStyle} />
      </Link.Anchor>
    </div>
  </div>
)

AsWrapper.story = {
  name: 'As wrapper',
}

export const Primary = () => (
  <div>
    <div>
      <span style={titleStyle}>link:</span>
      <Link primary to="/">
        Show more
      </Link>
    </div>
    <div>
      <span style={titleStyle}>anchor link:</span>
      <Link.Anchor primary href="/">
        Show more
      </Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>external anchor link:</span>
      <Link.Anchor primary external href="/">
        More information
      </Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>doc link:</span>
      <Link.DocLink primary path="/getting-started">
        Getting Started
      </Link.DocLink>
    </div>
  </div>
)

export const Secondary = () => (
  <div>
    <div>
      <span style={titleStyle}>link:</span>
      <Link secondary to="/">
        Show more
      </Link>
    </div>
    <div>
      <span style={titleStyle}>anchor link:</span>
      <Link.Anchor secondary href="/">
        Show more
      </Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>external anchor link:</span>
      <Link.Anchor secondary external href="/">
        More information
      </Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>doc link:</span>
      <Link.DocLink secondary path="/getting-started">
        Getting Started
      </Link.DocLink>
    </div>
  </div>
)

export const ShowVisited = () => (
  <div>
    <div>
      <span style={titleStyle}>link:</span>
      <Link primary showVisited to="/">
        Show more
      </Link>
    </div>
    <div>
      <span style={titleStyle}>anchor link:</span>
      <Link.Anchor primary showVisited href="/">
        Show more
      </Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>external anchor link:</span>
      <Link.Anchor external primary showVisited href="/">
        More information
      </Link.Anchor>
    </div>
    <div>
      <span style={titleStyle}>doc link:</span>
      <Link.DocLink primary showVisited path="/getting-started">
        Getting Started
      </Link.DocLink>
    </div>
  </div>
)

export const Disabled = () => (
  <div>
    <div>
      <span style={titleStyle}>link:</span>
      <Link primary disabled to="/">
        Show more
      </Link>
    </div>
    <div>
      <span style={titleStyle}>anchor link:</span>
      <Link.Anchor primary disabled href="/">
        Show more
      </Link.Anchor>
    </div>
  </div>
)
