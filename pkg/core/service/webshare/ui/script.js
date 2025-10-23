document.addEventListener('DOMContentLoaded', () => {
    const views = {
        login: document.getElementById('login-view'),
        files: document.getElementById('files-view'),
    };

    const loginForm = document.getElementById('login-form');
    const passcode_input = document.getElementById('passcode-input');
    const loginError = document.getElementById('login-error');

    const fileList = document.getElementById('file-list');
    const emptyListMessage = document.getElementById('empty-list-message');

    const uploadForm = document.getElementById('upload-form');
    const uploadInput = document.getElementById('upload-input');
    const uploadStatus = document.getElementById('upload-status');

    function showView(viewName) {
        Object.values(views).forEach(view => view.classList.remove('active'));
        views[viewName].classList.add('active');
    }

    async function fetchFiles() {
        try {
            const response = await fetch('/api/shared-files');
            if (!response.ok) {
                if (response.status === 401) {
                    showView('login');
                }
                return;
            }
            const files = await response.json();
            renderFiles(files);
            showView('files');
        } catch (error) {
            console.error('Error fetching files:', error);
            showView('login');
        }
    }

    function renderFiles(files) {
        fileList.innerHTML = '';
        if (!files || files.length === 0) {
            emptyListMessage.style.display = 'block';
        } else {
            emptyListMessage.style.display = 'none';
            files.forEach(file => {
                const li = document.createElement('li');
                li.className = 'file-item';
                li.innerHTML = `
                    <div class="file-info">
                        <span class="file-name">${escapeHTML(file.name)}</span>
                        <span class="file-size">${formatBytes(file.size)}</span>
                    </div>
                    <a href="/api/download?uuid=${file.uuid}" class="download-button" download>Download</a>
                `;
                fileList.appendChild(li);
            });
        }
    }

    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const passcode = passcode_input.value;
        loginError.textContent = '';

        try {
            const formData = new FormData();
            formData.append('passcode', passcode);

            const response = await fetch('/api/login', {
                method: 'POST',
                body: new URLSearchParams(formData)
            });

            if (response.ok) {
                await fetchFiles();
            } else {
                loginError.textContent = 'Invalid passcode. Please try again.';
            }
        } catch (error) {
            console.error('Login error:', error);
            loginError.textContent = 'An error occurred. Please try again.';
        }
    });

    uploadForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const file = uploadInput.files[0];
        if (!file) {
            uploadStatus.textContent = 'Please select a file to upload.';
            uploadStatus.style.color = '#fa383e';
            return;
        }

        uploadStatus.textContent = 'Uploading...';
        uploadStatus.style.color = '#1877f2';

        const formData = new FormData();
        formData.append('file', file);

        try {
            const response = await fetch('/api/upload', {
                method: 'POST',
                body: formData,
            });

            if (response.ok) {
                uploadStatus.textContent = 'File uploaded successfully!';
                uploadStatus.style.color = '#42b72a';
                uploadInput.value = ''; // Clear the input
                await fetchFiles(); // Refresh the file list
            } else {
                const errorText = await response.text();
                throw new Error(errorText || 'Upload failed');
            }
        } catch (error) {
            console.error('Upload error:', error);
            uploadStatus.textContent = `Error: ${error.message}`;
            uploadStatus.style.color = '#fa383e';
        }
    });


    function formatBytes(bytes, decimals = 2) {
        if (bytes === 0) return '0 Bytes';
        const k = 1024;
        const dm = decimals < 0 ? 0 : decimals;
        const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
        const i = Math.floor(Math.log(bytes) / Math.log(k));
        return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
    }

    function escapeHTML(str) {
        const p = document.createElement('p');
        p.appendChild(document.createTextNode(str));
        return p.innerHTML;
    }

    // Initial check
    fetchFiles();
});
