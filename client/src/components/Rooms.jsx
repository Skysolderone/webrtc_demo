import { useEffect, useRef } from 'react';
import { useLocation } from 'react-router-dom';
const Room = () => {
  const location = useLocation();
  const userVideo = useRef();
  const userStream = useRef();
  const partnerVideo = useRef();
  const peerRef = useRef();
  const webSocketRef = useRef();
  const openCamera = async () => {
    const constraints = {
      video: true,
      audio: true,
    };
    navigator.mediaDevices.getUserMedia(constraints).then((stream)=>{
        userVideo.current.srcObject=stream
        userStream.current=stream
    })
  };
  useEffect(()=>{
    
  })
};

export default Room;
