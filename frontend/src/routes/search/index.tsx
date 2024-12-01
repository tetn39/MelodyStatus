import { createFileRoute } from '@tanstack/react-router'
import { useTracks } from './-api/useSearchSpotify'
import('@/index.css')

export const Route = createFileRoute('/search/')({
	component: RouteComponent,
})

function RouteComponent() {
	const { data, isPending, error } = useTracks()

	if (isPending) return 'Loading...'

	if (error) return `An error has occurred: ${error.message}`

	return (
		<div>
			<header>
				<p>MelodyStatus v2</p>
				<a href="https://reactjs.org" target="_blank" rel="noopener noreferrer">
					Learn React
				</a>
				<div className="text-5xl">{data.name}</div>
				<div className="text-5xl">{data.artists[0].name}</div>
			</header>
		</div>
	)
}
