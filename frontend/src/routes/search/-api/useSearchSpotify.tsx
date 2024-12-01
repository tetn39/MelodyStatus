import axios from 'axios'
import { useQuery } from '@tanstack/react-query'

const clientId = import.meta.env.VITE_SPOTIFY_CLIENT_ID
const clientSecret = import.meta.env.VITE_SPOTIFY_CLIENT_SECRET

const fetchSpotifyToken = async (): Promise<string> => {
	const encodedCredentials = btoa(`${clientId}:${clientSecret}`)

	const response = await axios({
		method: 'post',
		url: 'https://accounts.spotify.com/api/token',
		data: 'grant_type=client_credentials',
		headers: {
			'Content-Type': 'application/x-www-form-urlencoded',
			Authorization: `Basic ${encodedCredentials}`,
		},
	})

	return response.data.access_token
}

const fetchTracks = async (accessToken: string) => {
	const response = await axios.get(
		'https://api.spotify.com/v1/tracks/78MS9OT9Flm2V78fhTw4rk',
		{
			headers: {
				Authorization: `Bearer ${accessToken}`,
			},
		},
	)
	console.log(response.data)
	return response.data
}

export const useSpotifyToken = () => {
	return useQuery({
		queryKey: ['spotifyToken'],
		queryFn: fetchSpotifyToken,
		staleTime: 1000 * 60 * 60,
	})
}

export const useTracks = () => {
	const { data: accessToken } = useSpotifyToken()
	return useQuery({
		queryKey: ['tracks'],
		queryFn: () => fetchTracks(accessToken || ''),
		enabled: !!accessToken,
	})
}
