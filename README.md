<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">

</head>
<body>

<div>


<header>
<h1>

<img src="Logo.png" alt="SudoSpot Logo" width=90 height= 90  style="vertical-align: middle;">
SudoSpot: Spotify Terminal Player
</h1>
</header>

<p>
SudoSpot is a minimalist, cross-platform Spotify client built for the command line. It uses the <a href="https://github.com/charmbracelet/bubbletea">Bubble Tea TUI framework</a> and the Spotify Web API to provide real-time status and playback control for any device connected to your Spotify account (phone, TV, desktop).
</p>


<section>
<h2>✨ Features</h2>
<ul>
<li><strong>Real-time Status:</strong> Displays the currently playing track, artist, and progress.</li>
<li><strong>Universal Control:</strong> Control music playing on <strong>any device</strong> linked to your Spotify account.</li>
<li><strong>Full Playback Control:</strong> Play, pause, skip, and seek directly from the terminal.</li>
<li><strong>Cool Visualizer Line:</strong> A stylized progress bar acts as a visualizer line, showing the current progress through the track.</li>
</ul>


<div>
<p>Terminal Screenshot Placeholder (Illustrates the UI)</p>

<img src="SudoSpot.png" alt="SudoSpot Terminal Screenshot" style="width: 100%; max-width: 600px; height: auto; border: 1px solid #1DB954;">
</div>
</section>


<section>
<h2>🚀 Installation</h2>


<h3>Installation</h3>

<h4>Option 1: Homebrew (Recommended)</h4>
<p>If you have Homebrew installed, you can install SudoSpot directly from the repository:</p>
<pre><code>brew tap Sudo-Aju/sudospot
brew install sudospot</code></pre>

<h4>Option 2: Build from Source</h4>
<p>1. Clone the Repository:</p>
<pre><code>git clone https://github.com/Sudo-Aju/sudospot.git
cd sudospot</code></pre>

<p>2. Build the Standalone Executable:</p>
<pre><code>go build -o sudospot cmd/sudospot/main.go</code></pre>

<p>3. Move to your PATH (optional):</p>
<pre><code>mv sudospot /usr/local/bin/</code></pre>

<h3>Configuration & Authentication</h3>
<p>SudoSpot stores your credentials and tokens locally so you only need to log in once.</p>

<p>1. Run <code>sudospot</code> in your terminal.</p>
<p>2. If it's your first time, you will be prompted to enter your Spotify Client ID and Secret.
(You can get these from the <a href="https://developer.spotify.com/dashboard/">Spotify Developer Dashboard</a>).
</p>
<p>3. A browser window will open to authorize the application.</p>
<p>4. Once authorized, your credentials and token are saved to your user config directory (e.g., <code>~/.config/sudospot/</code>), and you won't need to do this again.</p>
</section>


<section>
<h2>🕹️ Controls</h2>
<p>SudoSpot allows you to control the active Spotify playback on any device connected to your account.</p>

<div>
<table>
<thead>
<tr>
<th>Key</th>
<th>Action</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<th>Space</th>
<td>Play / Pause Toggle</td>
<td>Pauses the currently playing track or resumes playback.</td>
</tr>
<tr>
<th>N</th>
<td>Next Track</td>
<td>Skips to the next song in the queue.</td>
</tr>
<tr>
<th>P</th>
<td>Previous Track</td>
<td>Skips back to the previous song.</td>
</tr>
<tr>
<th>Right Arrow</th>
<td>Seek Forward (10s)</td>
<td>Jumps the playback 10 seconds ahead in the track.</td>
</tr>
<tr>
<th>Left Arrow</th>
<td>Seek Backward (10s)</td>
<td>Jumps the playback 10 seconds back in the track.</td>
</tr>
<tr>
<th>Q / Ctrl+C</th>
<td>Quit</td>
<td>Exits the SudoSpot application gracefully.</td>
</tr>
</tbody>
</table>
</div>
</section>


<section>
<div>
<h2>🤝 Attribution</h2>
<p>SudoSpot is built on the shoulders of these fantastic Go libraries:</p>
<ul>
<li><a href="https://github.com/charmbracelet/bubbletea">Bubble Tea</a>: The powerful TUI framework.</li>
<li><a href="https://github.com/zmb3/spotify/v2">go-spotify-client</a>: The robust Spotify Web API wrapper.</li>
</ul>
</div>
<div>
<h2>👨‍💻 Contributing</h2>
<p>
Contributions are welcome! Feel free to open an issue or submit a pull request if you have ideas for improvements or new features.
</p>
</div>
</section>

<footer>
SudoSpot Project Documentation
</footer>

</div>
</body>
</html>
