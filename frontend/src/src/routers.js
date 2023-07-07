import React, {useContext} from "react";
import {Navigate, Routes, Route, BrowserRouter as Router, Outlet} from "react-router-dom";

import Buildings from "./pages/buildings";
import Elevator from "./pages/elevator";
import Elevators from "./pages/elevators";
import Screens from "./pages/screeens";
import Zhk from "./pages/zhk";
import Admin from "./pages/admin";

// import LogoutPage from "./pages/logout"
// import ChatPage from "./pages/chat"

// import AuthContext from "./authContext/Context";

// function ProtectedRoute({children}) {
//     const {accessToken} = useContext(AuthContext)
//     if (!accessToken) {
//         return <Navigate to={'/auth/login'} replace/>
//     }
//
//     return children ? children : <Outlet/>
// }

function AppRoutes() {
    return (
        <Router>
            <Routes>
                <Route path="/admin" element={<Outlet/>}>
                    <Route path="" element={<Admin/>}/>
                    <Route path=":screen_id" element={<Admin/>}/>
                    <Route path="*" element={<Navigate to={'/admin/'} replace/>}/>
                </Route>
                <Route path="/screen/:screen_id" element={<Elevator/>}/>
                <Route path="*" element={<Navigate to={'/admin/'} replace/>}/>
            </Routes>
        </Router>
    );
}

export default AppRoutes