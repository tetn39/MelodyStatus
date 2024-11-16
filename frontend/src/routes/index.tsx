import { createFileRoute } from '@tanstack/react-router'
import { useHelloWorld } from '@/routes/-api/useHelloWorld'
import { Button } from '@/components/ui/button'
import('@/index.css')

export const Route = createFileRoute('/')({
	component: RouteComponent,
})

function RouteComponent() {
	// localhost:8080/api/v1/helloからデータを取得する
	const { data, isPending, error } = useHelloWorld()

	if (isPending) return 'Loading...'

	if (error) return `An error has occurred: ${error.message}`

	return (
		<div>
			<header>
				<p>MelodyStatus v2</p>
				<a href="https://reactjs.org" target="_blank" rel="noopener noreferrer">
					Learn React
				</a>
				<div className="text-5xl">{data.message}</div>
				<Button>Click me</Button>
			</header>
		</div>
	)
}
