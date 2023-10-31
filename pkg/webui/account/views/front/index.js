

import React from 'react'
import { Routes, Route, Navigate, useLocation } from 'react-router-dom'

import authRoutes from '@account/constants/auth-routes'

import Footer from '@ttn-lw/containers/footer'

import Header from '@account/containers/header'

import Login from '@account/views/login'
import TokenLogin from '@account/views/token-login'
import CreateAccount from '@account/views/create-account'
import ForgotPassword from '@account/views/forgot-password'
import UpdatePassword from '@account/views/update-password'
import FrontNotFound from '@account/views/front-not-found'
import Validate from '@account/views/validate'

import { selectApplicationRootPath } from '@ttn-lw/lib/selectors/env'

import style from './front.styl'

const FrontView = () => {
  const location = useLocation()

  return (
    <div className={style.container}>
      <section className={style.content}>
        <Header />
        <div className={style.main}>
          <Routes>
            <Route path="/login" Component={Login} />
            <Route path="/token-login" Component={TokenLogin} />
            <Route path="/register" Component={CreateAccount} />
            <Route path="/forgot-password" Component={ForgotPassword} />
            <Route path="/update-password" Component={UpdatePassword} />
            <Route path="/validate" Component={Validate} />
            <Route index element={<Navigate to="/login" />} />
            {authRoutes.map(({ path }) => (
              <Route
                path={path}
                key={path}
                element={
                  <Navigate to={`/login?n=${selectApplicationRootPath()}${location.pathname}`} />
                }
              />
            ))}
            <Route Component={FrontNotFound} />
          </Routes>
        </div>
      </section>
      <section className={style.visual} />
      <Footer className={style.footer} transparent />
    </div>
  )
}

export default FrontView
