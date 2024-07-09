import React from 'react';
import { useNavigate } from 'react-router-dom';
const CreateRoom = () => {
  const nav = useNavigate();
  const create = async (e) => {
    e.preventDefault();
    const resp = await fetch('http://localhost:8080/create');
    const { room_id } = await resp.json();
    console.log(`room id${room_id}`);
    nav(`/room/${room_id}`, { state: { id: room_id } });
  };
  return (
    <div>
      <button onClick={create}>Create room</button>
    </div>
  );
};

export default CreateRoom;
