import "./action.css";

const Action = ({ todo }) => {
  return (
    <div className="action">
      <button className="dots-container">
        <div className="dot"></div>
        <div className="dot"></div>
        <div className="dot"></div>
      </button>
      <div className="action-items">
        {todo.map((item, index) => {
          return (
            <div key={index} className="action-item">
              {item}
            </div>
          );
        })}
        <div className="triangle"></div>
      </div>
    </div>
  );
};

export default Action;
