import * as htmlToImage from 'html-to-image';
import { toPng, toJpeg, toBlob, toPixelData, toSvg } from 'html-to-image';
import Save from './Save'

function GetImg(node: any) {
htmlToImage.toPng(node)
  .then(function (dataUrl) {
    var img = new Image();
    img.src = dataUrl;
    var a = document.createElement("a"); //Create <a>
    a.href = img.src //Image Base64 Goes here
    a.download = "Image.png"; //File name Here
    a.click(); //Downloaded file   
   document.body.removeChild(a);
   URL.revokeObjectURL(img.src);  
})
  .catch(function (error) {
    console.error('oops, something went wrong!', error);
  });
}

export default GetImg
