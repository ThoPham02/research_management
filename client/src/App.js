import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { useSelector } from "react-redux";
import React from "react";

import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import { publicRoutes, privateRoutes } from "./Router";
import { roleSelector } from "./store/selectors";
import { DefaultLayout } from "./components/Layout";

function App() {
  const role = useSelector(roleSelector);
  
  return (
    <Router>
      <Routes>
        {publicRoutes.map((route, index) => {
          const Layout = route.Layout || DefaultLayout;
          const Page = route.component;
          return (
            <Route
              key={index}
              path={route.path}
              element={
                <Layout>
                  <Page />
                </Layout>
              }
            />
          );
        })}
        {privateRoutes
          .filter((route) => route.role === role)
          .map((route, index) => {
            const Layout = route.Layout || DefaultLayout;
            const Page = route.component;
            return (
              <Route
                key={index}
                path={route.path}
                element={
                  <Layout>
                    <Page />
                  </Layout>
                }
              />
            );
          })}
      </Routes>
    </Router>
  );
}

export default App;
