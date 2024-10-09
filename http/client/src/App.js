
import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
    const [originalImage, setOriginalImage] = useState(null);
    const [processedImage, setProcessedImage] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');


  const triggerFileSelect = () => {

    document.getElementById('file-input').click();

  }
    
    const handleFileChange = (event) => {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = (e) => {
                setOriginalImage(e.target.result);
            };
            reader.readAsDataURL(file);
            uploadImage(file);
        }
    };

    const handleDrop = (event) => {
        event.preventDefault();
        const file = event.dataTransfer.files[0];
        if (file) {
            handleFileChange({ target: { files: [file] } });
        }
    };

    const uploadImage = async (file) => {
        const formData = new FormData();
        formData.append('image', file);

        setLoading(true);
        setError('');
        try {
      const response = await axios.post('http://localhost:10000/api/bgremove', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },

                responseType: 'blob',
            });

          const imageUrl = URL.createObjectURL(response.data);
            setProcessedImage(imageUrl);
        } catch (error) {
            setError('Error uploading image. Please try again.');
            console.error('Error uploading image:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="App">
            <header className="App-header">
                <h1>Background Remover</h1>
                <p>Remove image backgrounds easily!</p>
            </header>
            <main className="upload-section">
                <h2>Upload Your Image</h2>
                <div
                    className={`upload-area ${loading ? 'loading' : ''}`}
                    onDrop={handleDrop}
                    onDragOver={(e) => e.preventDefault()}
                    onClick={triggerFileSelect}  
              >
                    {loading ? (
                        <div className="loader">Processing...</div>
                    ) : (
                        <>
                            <p>Drag & drop your image here or click to upload</p>
                            <input id="file-input" type="file" accept="image/*" onChange={handleFileChange} />
                        </>
                    )}
                </div>
                {error && <p className="error">{error}</p>}
                {originalImage && (
                    <div className="image-preview">
                        <h4>Original Image</h4>
                        <img src={originalImage} alt="Original" />
                    </div>
                )}
                {processedImage && (
                    <div className="image-preview">
                        <h4>Processed Image</h4>
                        <img src={processedImage} alt="Processed" />
                        <a href={processedImage} download className="btn">Download Processed Image</a>
                    </div>
                )}
            </main>
        </div>
    );
}

export default App;
