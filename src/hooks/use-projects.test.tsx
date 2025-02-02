import { renderHook, waitFor } from '@testing-library/react'

import { Status } from '$models/pipeline'
import { createProject } from '$test/objects'
import { createWrapper } from '$test/react-query'
import { useProjects } from './use-projects'

describe('useProjects', () => {
  it('should fetch projects', async () => {
    const groupId = 12

    const project = createProject('project-1')
    const status = Status.RUNNING

    global.fetch = vi.fn().mockResolvedValueOnce({
      json: () => Promise.resolve({ [status]: [project] })
    })

    const { result } = renderHook(() => useProjects(groupId), {
      wrapper: createWrapper()
    })

    await waitFor(() => expect(result.current.isSuccess).toBe(true))

    expect(result.current.data?.get(status)).toEqual([project])
    expect(global.fetch).toHaveBeenCalledWith(
      expect.stringContaining(`/api/groups/${groupId}/projects`)
    )
  })

  it('should return empty map when groupId is undefined', async () => {
    const groupId = undefined

    global.fetch = vi.fn().mockResolvedValueOnce({
      json: () => Promise.resolve()
    })

    const { result } = renderHook(() => useProjects(groupId), {
      wrapper: createWrapper()
    })

    await waitFor(() => expect(result.current.isSuccess).toBe(true))

    expect(global.fetch).not.toHaveBeenCalled()
    expect(result.current.data).toEqual(new Map())
  })
})
