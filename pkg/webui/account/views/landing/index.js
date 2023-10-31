

import React from 'react'
import { Routes, Route } from 'react-router-dom'
import classnames from 'classnames'

import authRoutes from '@account/constants/auth-routes'

import sidebarStyle from '@ttn-lw/components/navigation/side/side.styl'

import Footer from '@ttn-lw/containers/footer'

import { FullViewErrorInner } from '@ttn-lw/lib/components/full-view-error'

import Header from '@account/containers/header'

import style from './landing.styl'

const GenericNotFound = () => <FullViewErrorInner error={{ statusCode: 404 }} />

const Landing = () => (
  <div className={style.container}>
    <Header />
    <main className={style.main}>
      <div className={style.stage} id="stage">
        <div id="sidebar" className={sidebarStyle.container} />
        <div className={classnames('breadcrumbs', style.desktopBreadcrumbs)} />
        <Routes>
          {authRoutes.map(route => (
            <Route {...route} key={route.path} />
          ))}
          <Route path="*" element={<GenericNotFound />} />
        </Routes>
      </div>
    </main>
    <Footer />
  </div>
)

export default Landing
