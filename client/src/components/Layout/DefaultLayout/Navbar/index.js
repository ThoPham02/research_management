import "./navbar.css";
import { loginSelector } from "../../../../store/selectors";

import { Link } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { BiUserCircle } from "react-icons/bi";
import loginSlice from "../../../../pages/Login/LoginSlice";

const NavBar = () => {
  const login = useSelector(loginSelector);

  const dispatch = useDispatch();
  function handleLogout() {
    dispatch(loginSlice.actions.logout());
  }
  let switchLink = ""
  if (login.user.type) {
    switchLink = "/" + login.user.type
  }

  var loginAccount = (
    <div>
      <Link to="login">
        <BiUserCircle />
        <span>Đăng nhập</span>
      </Link>
    </div>
  );

  if (login.user.name !== undefined) {
    loginAccount = (
      <div>
        <span>Xin Chào, {login.user.name}</span>
        <ul className="account__nav">
          <li>
            <Link to="/user-info">Thông tin</Link>
          </li>
          <li onClick={handleLogout} to="/login">
            Đăng xuất
          </li>
        </ul>
      </div>
    );
  }

  return (
    <ul className="nav">
      <li>
        <Link to="/">Trang Chủ</Link>
      </li>
      <li>
        <Link to={`${switchLink}/topic`}>Đề tài</Link>
      </li>
      <li>
        <Link to={`${switchLink}/statistical`}>Thống Kê</Link>
      </li>
      <li>
        <Link to="/category">Danh mục</Link>
      </li>
      <li>
        <Link to="/contact">Liên Hệ</Link>
      </li>
      <li className="nav__login">
        {loginAccount}
      </li>
    </ul>
  );
};

export default NavBar;