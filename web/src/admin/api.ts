import api from '../../src/services/api'

export interface Applicant {
    id?: number
    applicant_type: 'student' | 'employee'
    name: string
    phone: string
    email: string
    purposes?: string
    notes?: string
    student: {
        school: string
        major: string
        enrollmentYear: string
        generateYear?: string
    }
    employee: {
        company: string
        position: string
        industry: string
        workExperience: number
    }
}

// 获取数据列表
export const fetchDataList = async (): Promise<Applicant[]> => {
    const response = await api.get('/admin/datalist/')
    return response.data
}

// 创建数据
export const createData = async (data: Applicant): Promise<Applicant> => {
    const response = await api.post('/submit', data)
    return response.data
}

// 更新数据
export const updateData = async (id: number, data: Applicant): Promise<Applicant> => {
    const response = await api.put(`/admin/edit/${id}`, data)
    return response.data
}

// 删除数据
export const deleteData = async (id: number): Promise<void> => {
    await api.delete(`/admin/data/${id}`)
}