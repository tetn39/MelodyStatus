import { StrictMode } from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider, createRouter } from '@tanstack/react-router'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { routeTree } from './routeTree.gen'

const queryClient = new QueryClient()

const router = createRouter({
	routeTree,
	context: {
		queryClient,
	},
	// リンクにホバーした時点でプリロードする
	defaultPreload: 'intent',
	// キャッシュが常に古いものとして、データを再取得されるようにする（キャッシュはtanstack-queryで管理するため）
	defaultPreloadStaleTime: 0,
})

declare module '@tanstack/react-router' {
	interface Register {
		router: typeof router
	}
}

const rootElement = document.getElementById('root')

if (rootElement && !rootElement.innerHTML) {
	const root = ReactDOM.createRoot(rootElement)
	root.render(
		<StrictMode>
			<QueryClientProvider client={queryClient}>
				<RouterProvider router={router} />
			</QueryClientProvider>
		</StrictMode>,
	)
} else {
	console.error('ルート要素が見つかりませんでした')
}
