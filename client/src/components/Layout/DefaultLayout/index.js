import Header from "./Header";
import Footer from "./Footer";
import NavBar from "./Navbar";

const background = require("./background.png")

const DefaultLayout = ({ children }) => {
  return (
    <div style={{backgroundImage: `url(${background})`, backgroundSize: "cover"}}>
      <Header />
      <NavBar />
      <div className="container" style={{ 
        minHeight: "calc(100vh - 219px)",
        marginTop: "4px",
      }}>
        {children}
      </div>
      <Footer />
    </ div>
  );
};

export default DefaultLayout;
