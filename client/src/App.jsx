import { BroswerRouter, Route, Routes } from 'react-router-dom';
// import './App.css';
import CreateRoom from './components/CreateRoom';
import Room from './components/Rooms';
function App() {
  return (
    <div className="App">
      <BroswerRouter>
        <Routes>
          <Route path="/" element={<CreateRoom />}></Route>
          <Route path="/room/:roomID" element={<Room />}></Route>
        </Routes>
      </BroswerRouter>
    </div>
  );
}

export default App;
