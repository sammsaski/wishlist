import './App.css';
import DisplayItems from './components/DisplayItems';
import Actions from './components/Actions';

function App() {
  return (
    <div>
      <div className="container-fluid">
        <div className="row">
          <div className="col-1"></div>
          <div className="col-10 text-center">
            <div className="home-text">Create a wishlist!</div>
          </div>
          <div className="col-1"></div>
        </div>
      </div>

      <div className="container-fluid">
        <div className="row">
          <div className="col-12 text-center">
            <div className="row d-flex justify-content-center">
              <div className="col-6">
                <DisplayItems></DisplayItems>
              </div>
            </div>
          </div>
        </div>

        <div className="container-fluid">
          <div className="row">
            <div className="col-12 text-center">
              <div className="row d-flex justify-content-center">
                <Actions></Actions>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
