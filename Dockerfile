# Step 1: Use an official Golang image as a base
FROM golang:1.20-alpine AS builder

# Step 2: Install dependencies for Alpine
RUN apk add --no-cache \
    gcc \
    g++ \
    make \
    ffmpeg \
    git \
    nodejs \
    npm \
    curl \
    bash

# Step 3: Install Miniconda
# Download and install Miniconda for managing Python and conda packages
RUN curl -sSLo /tmp/miniconda.sh https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh && \
    bash /tmp/miniconda.sh -b -p /opt/conda && \
    rm /tmp/miniconda.sh

# Set conda path
ENV PATH=/opt/conda/bin:$PATH

# Step 4: Create a conda environment and install WhisperX
RUN conda create -n whisperx-env python=3.9 -y && \
    /opt/conda/bin/conda init bash && \
    echo "source activate whisperx-env" >> ~/.bashrc && \
    /bin/bash -c "source ~/.bashrc && conda install -c conda-forge whisperx -y"

# Step 5: Set the working directory
WORKDIR /app

# Step 6: Copy the Go application source code
COPY . .

# Step 7: Use npx to build Tailwind CSS
RUN npx tailwindcss build assets/css/tailwind.css -o assets/dist/styles.css

# Step 8: Build the Go application
RUN go mod tidy && go build -o myapp

# Step 9: Use a minimal image to reduce size
FROM alpine:latest

# Step 10: Install FFmpeg for audio processing
RUN apk add --no-cache ffmpeg

# Step 11: Copy the built Go app and assets from the builder stage
COPY --from=builder /app/myapp /usr/local/bin/myapp
COPY --from=builder /app/assets /app/assets

# Step 12: Set up Conda in the final image
COPY --from=builder /opt/conda /opt/conda
ENV PATH=/opt/conda/bin:$PATH

# Step 13: Set the default command
CMD ["myapp"]
