
  var video = document.getElementById('video');
  var canvas = document.getElementById('canvas');
  var context = canvas.getContext('2d');
  var vendorURL = window.URL;

  navigator.getMedia = navigator.getUserMedia;

  var config = { video: true, audio: false};
  const f1 = (stream) => {
    video.srcObject = stream;
    video.play();
  }
  const f2 = (error) => {
    console.log(error);
  }
  navigator.getMedia(config, f1, f2);

  document.getElementById("capture").addEventListener('click', function () {
      //execute this when u click anchor
      context.drawImage(video, 0, 0, canvas.width, canvas.height);
      //now save the image and set it in img tag
      var myImage = canvas.toDataURL("image/png"); 
      var imageElement = document.getElementById("img");  
      imageElement.src = myImage;
      document.getElementById("send").style.display = "inline";
  });